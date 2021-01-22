<?php

namespace App\RealWorld\Transformers;

class UserTransformer extends Transformer
{
    protected $resourceName = 'user';

    public function transform($data)
    {
        return [
            'id'        => $data['id'],
            'email'     => $data['email'],
            'token'     => $data['token'],
            'username'  => $data['username'],
            'bio'       => $data['bio'],
            'image'     => $data['image'],
            'karma'     => $data['karma'],
            'role'      => $data['role'],
            'tempkey'   => $data['tempkey']
        ];
    }
}