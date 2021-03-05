import React from "react";
import { useQuery } from "react-query";
import ValheimLogo from "./logo_valheim.png";
import "./StatusPage.css";

const ONE_MINUTE = 1000 * 60;

const fetchServerStatus = async () => {
  const response = await fetch("/api/server");
  if (!response.ok)
    throw new Error(`${response.status} ${response.statusText}`);

  return response.json();
};

export const StatusPage: React.FC = () => {
  const serverStatusQuery = useQuery("serverStatus", fetchServerStatus, {
    refetchInterval: ONE_MINUTE,
  });

  const isOnline = serverStatusQuery.data?.Name !== undefined;
  return (
    <div className="status-page-container">
      <img src={ValheimLogo} alt="logo valheim" />
      <div className="status-page-card">
        <div>{serverStatusQuery.data?.Name}</div>
        <div
          className="status-page-online-badge"
          style={{ backgroundColor: isOnline ? "#1e5a1b" : "#671313" }}
        >
          {isOnline ? "Online" : "Offline"}
        </div>
        <div>Players</div>
        <div>{serverStatusQuery.data?.Players?.Current}</div>
        <div>Version</div>
        <div>{serverStatusQuery.data?.Raw?.ExtraData?.Keywords}</div>
      </div>
    </div>
  );
};
