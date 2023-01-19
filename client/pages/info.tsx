import Layout from "../components/Layout";
import Button from "react-bootstrap/Button";

export default function Info() {
  return (
    <Layout>
      <div id='Info'>
        <h1>Which lang is on trend???</h1>
        <Button variant="primary" className="mt-3">Fetch 🐸</Button>
      </div>
    </Layout>
  );
};
