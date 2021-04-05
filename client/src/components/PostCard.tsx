import React from "react";
import { Link } from "react-router-dom";
import { PostType } from "src/types/PostType";
import { formatTimestamp } from "src/util/timeFormatter";

function PostCard({ post }: { post: PostType }) {
  return (
    <article className="grid grid-cols-12 mb-8">
      <div className="col-span-3 text-gray-500 font-mono">
        <span>{formatTimestamp(post.createdAt)}</span>
      </div>
      <div className="col-span-7">
        <h3 className="text-xl font-semibold mb-4 hover:underline">
          <Link to={`/post/${post.id}`}>{post.title}</Link>
        </h3>
        <p className="text-gray-500 my-4 text-truncate">{post.content}</p>
        <div className="inline-block">
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
      </div>
    </article>
  );
}

export default PostCard;
