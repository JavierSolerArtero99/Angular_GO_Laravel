<?php

namespace App\RealWorld\Transformers;

class ProductTransformer extends Transformer
{
    protected $resourceName = 'product';

    public function transform($data)
    {
        return [
            'name'              => $data['name'],
            'image'             => $data['image'],
            'price'             => $data['price'],
            'createdAt'         => $data['created_at']->toAtomString(),
            'updatedAt'         => $data['updated_at']->toAtomString(),
            'favoritesCount'    => $data['favoritesCount'],
            'author' => [
                'username'  => $data['user']['username'],
                'bio'       => $data['user']['bio'],
                'image'     => $data['user']['image'],
            ]
        ];
    }
}