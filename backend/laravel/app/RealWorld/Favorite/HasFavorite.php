<?php

namespace App\RealWorld\Favorite;

use App\Product;

trait HasFavorite
{
    /**
     * Favorite the given product.
     *
     * @param Product $product
     * @return mixed
     */
    public function favorite(Product $product)
    {
        if (! $this->hasFavorited($product))
        {
            return $this->favorites()->attach($product);
        }
    }

    /**
     * Unfavorite the given product.
     *
     * @param Product $product
     * @return mixed
     */
    public function unFavorite(Product $product)
    {
        return $this->favorites()->detach($product);
    }

    /**
     * Get the products favorited by the user.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsToMany
     */
    public function favorites()
    {
        return $this->belongsToMany(Product::class, 'favorites', 'user_id', 'product_id')->withTimestamps();
    }

    /**
     * Check if the user has favorited the given product.
     *
     * @param Product $product
     * @return bool
     */
    public function hasFavorited(Product $product)
    {
        return !! $this->favorites()->where('product_id', $product->id)->count();
    }
}