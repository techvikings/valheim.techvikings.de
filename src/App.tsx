import React from "react";
import "./App.css";
import { QueryClient, QueryClientProvider } from "react-query";
import { StatusPage } from "./StatusPage";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <StatusPage />
    </QueryClientProvider>
  );
}

export default App;
