import React from "react";

function PageNotFound() {
  return (
    <div className="w-full h-full flex flex-col items-center justify-center">
      <h1 className="text-xl">
        <strong>404</strong> | Page not found
      </h1>
      <a href="/" className="mt-4 hover:underline">
        ← Go back home
      </a>
    </div>
  );
}

export default PageNotFound;
