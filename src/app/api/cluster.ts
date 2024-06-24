import { NextApiRequest, NextApiResponse } from 'next';

let nodeIndex = 0;
const nodes = [
  'http://localhost:8081/cluster',
  'http://localhost:8082/cluster',
  'http://localhost:8083/cluster'
];

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  if (req.method === 'POST') {
    const response = await fetch(nodes[nodeIndex], {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
    });
    nodeIndex = (nodeIndex + 1) % nodes.length; // Simple round-robin load balancing
    const result = await response.json();
    res.status(200).json(result);
  } else {
    res.status(405).json({ message: 'Method Not Allowed' });
  }
  
}
