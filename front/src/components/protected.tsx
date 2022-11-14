import {Navigate} from "@tanstack/react-location";
import React from "react";

const Protected = ({ isLoggedIn, children }) => {
  if (!isLoggedIn) {
  return <Navigate to="/login" />;
  }
  return children;
};
export default Protected;