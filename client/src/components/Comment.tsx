import React from "react";
import { CommentType } from "src/types/CommentType";
import { formatTimestamp } from "src/util/timeFormatter";

function Comment({ comment }: { comment: CommentType }) {
  return (
    <article className="my-8">
      <div className="flex mb-2">
        <a href="#" className="flex items-center group">
          <img
            className="w-6 h-6 rounded-full mr-4"
            src="https://randomuser.me/api/portraits/men/82.jpg"
          />
          <span className="text-gray-500 group-hover:underline">
            Stuart Little
          </span>
        </a>
        <span className="text-gray-500 ml-auto">
          {formatTimestamp(comment.createdAt)}
        </span>
      </div>
      <p className="text-base text-gray-700">{comment.content}</p>
    </article>
  );
}

export default Comment;
