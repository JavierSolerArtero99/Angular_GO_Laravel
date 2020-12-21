<?php

namespace App\Http\Controllers\Api;

use Auth;
use App\User;
use App\Http\Requests\Api\LoginUser;
use App\Http\Requests\Api\RegisterUser;
use App\RealWorld\Transformers\UserTransformer;
use Illuminate\Support\Facades\Redis;

class AuthController extends ApiController
{
    /**
     * AuthController constructor.
     *
     * @param UserTransformer $transformer
     */
    public function __construct(UserTransformer $transformer)
    {
        $this->transformer = $transformer;
    }

    /**
     * Login user and return the user if successful.
     *
     * @param LoginUser $request
     * @return \Illuminate\Http\JsonResponse
     */
    public function login(LoginUser $request)
    {
        $credentials = $request->only('user.email', 'user.username');
        $credentials = $credentials['user'];
        $credentials['password'] = 'null';

        if ($data = Redis::get($credentials['username'])) {
            $json = json_decode($data);

            $user = \App\User::where('username', $json->{'username'})->first();
            if ($user['username'] === $user->getTempkey()[1] && $credentials['email'] === $user['email']) {
                $credentials['password'] = $user->getTempkey()[2];
            }
        };

        if (! Auth::once($credentials)) {
            return $this->respondFailedLogin();
        }

        $userauth = auth()->user();
        $user->setRememberToken($userauth);
        serialize($user);
        var_dump($user);
        return $this->respondWithTransformer($userauth);
    }

    /**
     * Register a new user and return the user if successful.
     *
     * @param RegisterUser $request
     * @return \Illuminate\Http\JsonResponse
     */
    public function register(RegisterUser $request)
    {
        $user = User::create([
            'username' => $request->input('user.username'),
            'email' => $request->input('user.email'),
            'password' => bcrypt($request->input('user.password')),
            'karma' => 0,
            'role' => 0,
        ]);

        return $this->respondWithTransformer($user);
    }
}
