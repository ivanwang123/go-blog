import axios from "axios";
import React, { useContext, useEffect, useState } from "react";
import { useQuery } from "react-query";
import Footer from "src/components/Footer";
import Navbar from "src/components/Navbar";
import PostCard from "src/components/PostCard";
import { PostType } from "src/types/PostType";

const fetchPosts = ({ queryKey }: any) => {
  const [_key, page] = queryKey;
  return axios.get(`/post/paginate?page=${page}`);
};

function HomePage() {
  const [page, setPage] = useState<number>(1);
  const {
    data: { data } = { data: null },
    error,
    isLoading,
    isError,
    isFetching,
    isPreviousData,
  } = useQuery<any, any>(["posts", page], fetchPosts, {
    keepPreviousData: true,
  });

  if (isLoading) return <h1>Loading...</h1>;
  if (isError) return <h1>Error getting posts</h1>;
  console.log("DATA", data);

  const pageForward = () => {
    if (data.hasMore) setPage(page + 1);
  };

  const pageBackward = () => {
    if (page > 1) setPage(page - 1);
    else setPage(1);
  };

  return (
    <div className="grid grid-cols-12 min-h-full">
      <Navbar />
      <main className="col-span-12 px-24">
        <h1 className="text-5xl font-semibold my-24">All posts</h1>
        <section className="post-feed flex flex-col">
          {data.posts.map((post: PostType) => (
            <PostCard post={post} key={post.id} />
          ))}
        </section>
        <div className="flex justify-center items-center font-mono mb-4">
          <button
            type="button"
            onClick={pageBackward}
            className="text-xl px-1 disabled:text-gray-400"
            disabled={page <= 1}
          >
            &lt;
          </button>
          <span className="mx-2">{page}</span>
          <button
            type="button"
            onClick={pageForward}
            className="text-xl px-1 disabled:text-gray-400"
            disabled={!data.hasMore}
          >
            &gt;
          </button>
        </div>
      </main>
      <Footer />
    </div>
  );
}

export default HomePage;
