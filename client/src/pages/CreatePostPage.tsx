import axios from "axios";
import React, { FormEvent, useContext } from "react";
import { useMutation } from "react-query";
import { useHistory } from "react-router";
import Footer from "src/components/Footer";
import Navbar from "src/components/Navbar";
import { useAuthContext } from "src/contexts/Auth";

const createPost = (formData: FormData) => {
  return axios.post("/post", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
};

function CreatePostPage() {
  const auth = useAuthContext();
  const history = useHistory();

  const { mutate, data, error, isError, isLoading, isSuccess } = useMutation<
    any,
    any,
    FormData
  >(createPost);

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const formData = new FormData(e.currentTarget);
    formData.append("user-id", auth.me!.id.toString());
    mutate(formData);
  };

  if (isSuccess) {
    history.push(`/post/${data.data.postId}`);
  }

  return (
    <div className="grid grid-cols-12 min-h-full">
      <Navbar />
      <main className="col-start-4 col-end-10 flex flex-col">
        <h1 className="text-4xl my-8">Create a post</h1>
        <section className="h-full">
          <form className="flex flex-col h-full mb-12" onSubmit={handleSubmit}>
            <label htmlFor="title" className="mb-2">
              Title
            </label>
            <input
              type="text"
              name="title"
              id="title"
              className="border border-gray-300 px-2 py-1 rounded"
              required
            />
            <label htmlFor="content" className="my-2">
              Content
            </label>
            <textarea
              name="content"
              id="content"
              className="h-96 border border-gray-300 px-2 py-1 rounded"
              required
            ></textarea>
            <button
              type="submit"
              className="bg-gray-100 border border-gray-300 py-1 mt-4 rounded"
            >
              {isLoading ? "Submitting..." : "Submit"}
            </button>
            <div className="text-red-500 mt-2">
              {isError && <p>{error.response.data.message}</p>}
            </div>
          </form>
        </section>
      </main>
      <Footer />
    </div>
  );
}

export default CreatePostPage;
