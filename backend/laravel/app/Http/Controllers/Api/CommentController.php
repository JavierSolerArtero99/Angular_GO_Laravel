<?php

namespace App\Http\Controllers\Api;

use App\Product;
use App\Comment;
use App\Http\Requests\Api\CreateComment;
use App\Http\Requests\Api\DeleteComment;
use App\RealWorld\Transformers\CommentTransformer;

class CommentController extends ApiController
{
    /**
     * CommentController constructor.
     *
     * @param CommentTransformer $transformer
     */
    public function __construct(CommentTransformer $transformer)
    {
        $this->transformer = $transformer;

        $this->middleware('auth.api')->except('index');
        $this->middleware('auth.api:optional')->only('index');
    }

    /**
     * Get all the comments of the product given by its slug.
     *
     * @param Product $product
     * @return \Illuminate\Http\JsonResponse
     */
    public function index(Product $product)
    {
        $comments = $product->comments()->get();

        return $this->respondWithTransformer($comments);
    }

    /**
     * Add a comment to the product given by its slug and return the comment if successful.
     *
     * @param CreateComment $request
     * @param Product $product
     * @return \Illuminate\Http\JsonResponse
     */
    public function store(CreateComment $request, Product $product)
    {
        $comment = $product->comments()->create([
            'body' => $request->input('comment.body'),
            'user_id' => auth()->id(),
        ]);

        return $this->respondWithTransformer($comment);
    }

    /**
     * Delete the comment given by its id.
     *
     * @param DeleteComment $request
     * @param $product
     * @param Comment $comment
     * @return \Illuminate\Http\JsonResponse
     */
    public function destroy(DeleteComment $request, $product, Comment $comment)
    {
        $comment->delete();

        return $this->respondSuccess();
    }
}
