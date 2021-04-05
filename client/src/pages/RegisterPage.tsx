import axios from "axios";
import React, { FormEvent } from "react";
import { useMutation } from "react-query";
import { Link, useHistory } from "react-router-dom";

const register = (formData: FormData) => {
  return axios.post("/user/register", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
};

function RegisterPage() {
  const history = useHistory();

  const { mutate, error, isError, isLoading, isSuccess } = useMutation<
    any,
    any,
    FormData
  >(register);

  if (isSuccess) {
    history.push("/login");
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
        <h1 className="text-2xl mb-8 mt-4">Create an account</h1>
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
            <label htmlFor="username" className="my-2">
              Username
            </label>
            <input
              type="text"
              name="username"
              id="username"
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
            <label htmlFor="confirm-password" className="my-2">
              Confirm Password
            </label>
            <input
              type="password"
              name="confirm-password"
              id="confirm-password"
              className="border border-gray-300 px-2 py-1 rounded"
              required
            />
            <button
              type="submit"
              className="bg-gray-100 border border-gray-300 py-1 mt-4 rounded"
            >
              Register
            </button>
            <div className="text-red-500 mt-2">
              {isError && <p>{error.response.data.message}</p>}
            </div>
          </form>
        </section>
        <div className="flex justify-center my-8">
          <p>
            Already have an account?{" "}
            <Link to="/login" className="underline">
              {isLoading ? "Registering..." : "Register"}
            </Link>
          </p>
        </div>
      </main>
    </div>
  );
}

export default RegisterPage;
