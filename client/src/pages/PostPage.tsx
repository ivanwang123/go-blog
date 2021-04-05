import axios from "axios";
import React, { FormEvent, useEffect, useState } from "react";
import {
  QueryObserverResult,
  RefetchOptions,
  useMutation,
  useQuery,
  UseQueryResult,
} from "react-query";
import { useParams } from "react-router-dom";
import Footer from "src/components/Footer";
import Navbar from "src/components/Navbar";
import Comment from "src/components/Comment";
import { formatTimestamp } from "src/util/timeFormatter";
import { CommentType } from "src/types/CommentType";
import { useAuthContext } from "src/contexts/Auth";

const fetchPosts = ({ queryKey }: any) => {
  const [_key, postId] = queryKey;
  return axios.get(`/post/${postId}`);
};

const fetchComments = ({ queryKey }: any) => {
  const [_key, postId] = queryKey;
  return axios.get(`/comment/post/${postId}`);
};

const createComment = (formData: FormData) => {
  return axios.post(`/comment`, formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
};

const userLikedPost = (userId: number, postId: number) => {
  return axios.get(`/like/${userId}/liked/${postId}`);
};

const toggleLike = (userId: number, postId: number) => {
  let formData = new FormData();
  formData.append("user-id", userId.toString());
  formData.append("post-id", postId.toString());

  return axios.post(`/like`, formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
};

function PostPage() {
  const { id }: any = useParams();
  const auth = useAuthContext();

  const {
    data: { data } = { data: null },
    error,
    isLoading,
    isError,
  } = useQuery<any, any>(["posts", id], fetchPosts);
  const comments = useQuery<any, any>(["comments", id], fetchComments);

  const [liked, setLiked] = useState<boolean>(false);

  useEffect(() => {
    window.scrollTo(0, 0);

    if (auth.isAuthenticated) {
      (async () => {
        try {
          const res = await userLikedPost(auth.me!.id, id);
          setLiked(res.data.liked);
        } catch (e) {
          setLiked(false);
        }
      })();
    }
  }, [auth.isAuthenticated]);

  const handleLike = () => {
    if (auth.isAuthenticated) {
      (async () => {
        try {
          const res = await toggleLike(auth.me!.id, id);
          setLiked(res.data.liked);
        } catch (e) {
          setLiked(false);
        }
      })();
    }
  };

  if (isLoading) return <h1>Loading...</h1>;
  if (isError) {
    return <h1>Error getting post</h1>;
  }

  return (
    <div className="grid grid-cols-12">
      <Navbar />
      <main className="col-span-12 px-4">
        <section className="grid grid-cols-12">
          <div className="col-start-3 col-end-10 border-b-2 border-dashed">
            <img
              className="w-full mt-6"
              src="https://github.blog/wp-content/uploads/2021/03/GitHub-Mobile-blog-hero.jpeg?w=2048"
            />
            <div className="relative post-content">
              <div className="font-mono text-gray-500 mt-24 mb-6">
                <span>{formatTimestamp(data.post.createdAt)}</span>
              </div>
              <h1 className="text-5xl font-semibold">{data.post.title}</h1>
              <div className="inline-block mt-6 mb-10">
                <a href="#" className="flex items-center group">
                  <img
                    className="w-8 h-8 rounded-full mr-4"
                    src="https://randomuser.me/api/portraits/men/82.jpg"
                  />
                  <span className="text-gray-500 group-hover:underline">
                    Stuart Little
                  </span>
                </a>
              </div>
              <p className="text-base text-gray-700 mb-12">
                {data.post.content}
              </p>
              <div className="flex items-center text-base text-gray-500 mb-6">
                Did you like this post?
                <button
                  type="button"
                  className="ml-2 focus:outline-none"
                  onClick={handleLike}
                >
                  {liked ? (
                    <img src="/res/liked.svg" />
                  ) : (
                    <img src="/res/like.svg" />
                  )}
                </button>
              </div>
            </div>
          </div>
          <div className="col-start-3 col-end-10">
            {comments.isError ||
            comments.isLoading ||
            !comments.data.data.comments.length ? (
              <>
                <h3 className="text-xl text-gray-500 my-6">No comments</h3>
                {auth.isAuthenticated && (
                  <SubmitComment postId={id} refetch={comments.refetch} />
                )}
              </>
            ) : (
              <>
                <h3 className="text-xl text-gray-500 mt-6">Comments</h3>
                {auth.isAuthenticated && (
                  <SubmitComment postId={id} refetch={comments.refetch} />
                )}
                {comments.data.data.comments.map((comment: CommentType) => (
                  <Comment comment={comment} />
                ))}
              </>
            )}
          </div>
        </section>
      </main>
      <Footer />
    </div>
  );
}

function SubmitComment({
  postId,
  refetch,
}: {
  postId: number;
  refetch: () => Promise<QueryObserverResult<any, any>>;
}) {
  const { mutate, isSuccess } = useMutation<any, any, FormData>(createComment);

  if (isSuccess) {
    refetch();
  }

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    let formData = new FormData(e.currentTarget);
    formData.append("post-id", postId.toString());
    mutate(formData);
  };

  return (
    <form className="flex flex-col mt-2 mb-8" onSubmit={handleSubmit}>
      <textarea
        name="content"
        id="content"
        className="h-24 border border-gray-300 px-2 py-1 rounded"
        placeholder="Write a comment..."
        required
      ></textarea>
      <button
        type="submit"
        className="bg-gray-100 border border-gray-300 py-1 mt-4 rounded"
      >
        Submit
      </button>
    </form>
  );
}

export default PostPage;
