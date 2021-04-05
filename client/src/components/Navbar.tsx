import axios from "axios";
import React, { useContext } from "react";
import { Link } from "react-router-dom";
import { useAuthContext } from "src/contexts/Auth";

function Navbar() {
  const auth = useAuthContext();

  const handleLogout = async () => {
    if (auth.isAuthenticated) {
      await axios.post("/user/logout");
    }
    auth.getMe();
  };

  return (
    <div className="col-span-12 h-20 grid grid-cols-12 place-items-center bg-gray-100">
      <h1 className="col-start-5 col-end-9 font-mono text-xl text-gray-800">
        <Link to="/">
          <div className="flex items-center">
            <div className="inline-block w-6 h-6 bg-gray-800 mr-4 rounded-full"></div>
            The Go Blog
          </div>
        </Link>
      </h1>
      <div className="col-start-9 col-end-13 font-mono text-gray-500 ml-auto mr-24">
        {auth.isAuthenticated ? (
          <>
            <a href="/create/post" className="mr-6 hover:underline">
              Write a post
            </a>
            <button
              type="button"
              className="hover:underline focus:outline-none"
              onClick={handleLogout}
            >
              Logout
            </button>
          </>
        ) : (
          <>
            <a href="/login" className="mr-6 hover:underline">
              Login
            </a>
            <a href="/register" className="hover:underline">
              Register
            </a>
          </>
        )}
      </div>
    </div>
  );
}

export default Navbar;
