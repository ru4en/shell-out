import React, { useEffect, useState } from 'react';
import axios from 'axios';

import { Card, Button } from 'react-bootstrap';

// Define the TemplateProps type
export type TemplateProps = {
  id: number;
  title: string;
  description: string;
  dockerfile: string;
};

// Define the component
const TemplateList: React.FC = () => {
  // State to hold the templates
  const [templates, setTemplates] = useState<TemplateProps[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  // Fetch the templates from the API
  useEffect(() => {
    const fetchTemplates = async () => {
      try {
        const response = await axios.get<TemplateProps[]>('/api/templates');
        setTemplates(response.data);
      } catch (err) {
        setError('Failed to fetch templates');
      } finally {
        setLoading(false);
      }
    };

    fetchTemplates();
  }, []);

  // Render the component
  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div>
      <h1>Templates</h1>
        {templates.map((template) => (
            <Card key={template.id} style={{ marginBottom: '10px' }}>
            <Card.Body>
                <Card.Title>{template.title}</Card.Title>
                <Card.Text>{template.description}</Card.Text>
                <Card.Text style={{ backgroundColor: '#f8f9fa', padding: '10px', borderRadius: '5px', textAlign: 'left' }}>
                    <Button 
                        className="btn btn-primary"
                        style={{ position: 'absolute', top: '10px', right: '10px' }}
                        onClick={() => {
                            navigator.clipboard.writeText(template.dockerfile);
                        }}
                    >Copy to clipboard</Button>
                    <pre>{template.dockerfile}</pre>
                </Card.Text>
            </Card.Body>
            </Card>
        ))}
    </div>
  );
};

export default TemplateList;
