import axios from "axios";
import React, { FormEvent, useContext } from "react";
import { useMutation } from "react-query";
import { Link, useHistory } from "react-router-dom";
import { useAuthContext } from "src/contexts/Auth";

const login = (formData: FormData) => {
  return axios.post("/user/login", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
};

function LoginPage() {
  const history = useHistory();
  const auth = useAuthContext();

  const { mutate, error, isError, isLoading, isSuccess } = useMutation<
    any,
    any,
    FormData
  >(login);

  if (isSuccess) {
    auth.getMe();
    history.push("/");
  }

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    mutate(new FormData(e.currentTarget));
  };

  return (
    <div className="grid grid-cols-12 min-h-full bg-gray-100">
      <main className="col-start-5 col-end-9 flex flex-col items-center">
        <Link to="/" className="mt-8">
          <div className="w-12 h-12 bg-black rounded-full"></div>
        </Link>
        <h1 className="text-2xl mb-8 mt-4">Log in to GoBlog</h1>
        <section className="w-full bg-white p-4 rounded">
          <form className="flex flex-col" onSubmit={handleSubmit}>
            <label htmlFor="email" className="mb-2">
              Email
            </label>
            <input
              type="email"
              name="email"
              id="email"
              className="border border-gray-300 px-2 py-1 rounded"
              required
            />
            <label htmlFor="password" className="my-2">
              Password
            </label>
            <input
              type="password"
              name="password"
              id="password"
              className="border border-gray-300 px-2 py-1 rounded"
              required
            />
            <button
              type="submit"
              className="bg-gray-100 border border-gray-300 py-1 mt-4 rounded"
            >
              {isLoading ? "Logging in..." : "Login"}
            </button>
            <div className="text-red-500 mt-2">
              {isError && <p>{error.response.data.message}</p>}
            </div>
          </form>
        </section>
        <div className="flex justify-center my-8">
          <p>
            New to GoBlog?{" "}
            <Link to="/register" className="underline">
              Create an account
            </Link>
          </p>
        </div>
      </main>
    </div>
  );
}

export default LoginPage;
