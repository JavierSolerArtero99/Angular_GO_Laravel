<?php

namespace App\Http\Controllers\Api;

use Auth;
use App\User;
use App\Http\Requests\Api\LoginUser;
use App\Http\Requests\Api\RegisterUser;
use App\RealWorld\Transformers\UserTransformer;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Redis;
use Illuminate\Support\Facades\Validator;

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
    public function login(Request $request)
    {

        $credentials = !$request['body'] ? $request->only('user.email', 'user.username') : $request['body'];

        $credentials = $credentials['user'];
        $credentials['password'] = 'null';

        $validator = Validator::make(
            $credentials,
            [
                'email' => 'required|email|max:255',
            ]
        );
        if ($validator->fails())
        {
            return $this->respondFailedLogin();
        }

        if ($data = Redis::get($credentials['username'])) {
            $json = json_decode($data);

            $user = \App\User::where('username', $json->{'username'})->first();
            if ($user['username'] === $user->getTempkey()[1] && $credentials['email'] === $user['email'] && $request->ip() === $user->getTempkey()[3]) {
                $credentials['password'] = $user->getTempkey()[2];
            }
        };

        if (! Auth::once($credentials)) {
            return $this->respondFailedLogin();
        }

        $userauth = auth()->user();
        $user->setRememberToken($userauth);
        serialize($user);
        return $this->respondWithTransformer($userauth);
    }

    /**
     * Return redis current users
     */
    public function currentUsers()
    {
        $result = array("users"=>Redis::get("current_users"));
        return $result;
    }

    /**
     * Return redis total users
     */
    public function totalUsers()
    {
        $result = array("users"=>Redis::get("total_users"));
        return $result;
    }

    public function hello() {
        return 'hello';
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
