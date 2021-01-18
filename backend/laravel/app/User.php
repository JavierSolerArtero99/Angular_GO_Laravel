<?php

namespace App;

use App\RealWorld\Follow\Followable;
use App\RealWorld\Favorite\HasFavorite;
use Illuminate\Notifications\Notifiable;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Tymon\JWTAuth\Contracts\JWTSubject;
use Tymon\JWTAuth\Facades\JWTAuth;

class User extends Authenticatable implements JWTSubject
{
    use Notifiable, Followable, HasFavorite;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'username', 'email', 'password', 'name', 'role', 'karma', 'tempkey'
    ];

    /**
     * The attributes that should be hidden for arrays.
     *
     * @var array
     */
    protected $hidden = [
        'password', 'remember_token',
    ];

    /**
     * Set the password using bcrypt hash.
     *
     * @param $value
     */
    public function setPasswordAttribute($value)
    {
        $this->attributes['password'] = (password_get_info($value)['algo'] === 0) ? bcrypt($value) : $value;
    }

    /**
     * Generate a JWT token for the user.
     *
     * @return string
     */
    public function getTokenAttribute()
    {
        return JWTAuth::fromUser($this);
    }

    /**
     * Get all the products by the user.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasMany
     */
    public function products()
    {
        return $this->hasMany(Article::class)->latest();
    }

    /**
     * Get all the comments by the user.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasMany
     */
    public function comments()
    {
        return $this->hasMany(Comment::class)->latest();
    }

    /**
     * Get all the products of the following users.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasMany
     */
    public function feed()
    {
        $followingIds = $this->following()->pluck('id')->toArray();

        return Article::loadRelations()->whereIn('user_id', $followingIds);
    }

    /**
     * Get the role Client(0)|Admin(1)
     *
     * @return boolean
     */
    public function isAdmin()
    {
        return 'role';
    }    
    
    /**
     * Get the key name for route model binding.
     *
     * @return string
     */
    public function getRouteKeyName()
    {
        return 'username';
    }

    /**
     * Get the tempkey
     *
     * @return $tempkey
     */
    public function getTempkey()
    {

        $encrypted = $this->attributes['tempkey'];
        if (!$encrypted) return null;
        
        $x = explode('#', $encrypted);

        $eUser = explode('@', $x[0]);
        $user = explode('%', $eUser[1]);
        $id = intVal($eUser[0]) / count($user);

        $username = '';
        foreach ($user as $key => $value) {
            $username .= chr($value/count($user));
        }

        $ePswd = explode('&', explode('!', $x[1])[0]);
        $pswd = '';
        foreach ($ePswd as $key => $value) {
            $pswd .= chr(sqrt(floatval($value)/$id));
        }

        $eIp = explode('$', explode('!', $x[1])[1]);
        $ip = '';
        foreach ($eIp as $key => $value) {
            $ip .= (sqrt(floatval($value))) . '.';
        }
        $ip = rtrim($ip, '. ');

        return [$id, $username, $pswd, $ip];
    }

    /**
     * Get the identifier that will be stored in the subject claim of the JWT.
     *
     * @return mixed
     */
    public function getJWTIdentifier()
    {
        return $this->getKey();
    }

    /**
     * Return a key value array, containing any custom claims to be added to the JWT.
     *
     * @return array
     */
    public function getJWTCustomClaims()
    {
        return [];
    }
}
