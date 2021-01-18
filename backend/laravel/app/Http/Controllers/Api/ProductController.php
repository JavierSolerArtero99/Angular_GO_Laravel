<?php

namespace App\Http\Controllers\Api;

use App\Tag;
use App\Product;
use App\RealWorld\Paginate\Paginate;
use App\RealWorld\Filters\ProductFilter;
use App\Http\Requests\Api\CreateProduct;
use App\Http\Requests\Api\UpdateProduct;
use App\Http\Requests\Api\DeleteProduct;
use App\RealWorld\Transformers\ProductTransformer;

class ProductController extends ApiController
{
    /**
     * ProductController constructor.
     *
     * @param ProductTransformer $transformer
     */
    public function __construct(ProductTransformer $transformer)
    {
        $this->transformer = $transformer;

        $this->middleware('auth.api')->except(['index', 'show']);
        $this->middleware('auth.api:optional')->only(['index', 'show']);
    }

    /**
     * Get all the products.
     *
     * @param ProductFilter $filter
     * @return \Illuminate\Http\JsonResponse
     */
    public function index(ProductFilter $filter)
    {
        $products = new Paginate(Product::loadRelations()->filter($filter));

        return $this->respondWithPagination($products);
    }

    /**
     * Create a new product and return the product if successful.
     *
     * @param CreateProduct $request
     * @return \Illuminate\Http\JsonResponse
     */
    public function store(CreateProduct $request)
    {
        $user = auth()->user();

        $product = $user->products()->create([
            'name' => $request->input('product.name'),
            'image' => $request->input('product.image'),
            'price' => $request->input('product.price'),
            'likes' => 0
        ]);

        return $this->respondWithTransformer($product);
    }

    /**
     * Get the product given by its slug.
     *
     * @param Product $product
     * @return \Illuminate\Http\JsonResponse
     */
    public function show(Product $product)
    {
        return $this->respondWithTransformer($product);
    }

    /**
     * Update the product given by its slug and return the product if successful.
     *
     * @param UpdateProduct $request
     * @param Product $product
     * @return \Illuminate\Http\JsonResponse
     */
    public function update(UpdateProduct $request, Product $product)
    {
        if ($request->has('product')) {
            $product->update($request->get('product'));
        }

        return $this->respondWithTransformer($product);
    }

    /**
     * Delete the product given by its slug.
     *
     * @param DeleteProduct $request
     * @param Product $product
     * @return \Illuminate\Http\JsonResponse
     */
    public function destroy(DeleteProduct $request, Product $product)
    {
        $product->delete();

        return $this->respondSuccess();
    }
}
