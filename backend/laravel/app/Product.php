<?php

namespace App;

use App\RealWorld\Slug\HasSlug;
use App\RealWorld\Filters\Filterable;
use App\RealWorld\Favorite\Favoritable;
use Illuminate\Database\Eloquent\Model;

class Article extends Model
{
    use Favoritable, Filterable, HasSlug;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'price', 'image'
    ];

    /**
     * Get the user that owns the article.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsTo
     */
    public function user()
    {
        return $this->belongsTo(User::class);
    }

    /**
     * Get the key name for route model binding.
     *
     * @return string
     */
    public function getRouteKeyName()
    {
        return 'slug';
    }

    /**
     * Get the attribute name to slugify.
     *
     * @return string
     */
    public function getSlugSourceColumn()
    {
        return 'title';
    }

    /**
     * Get list of values which are not allowed for this resource
     *
     * @return array
     */
    public function getBannedSlugValues()
    {
        return ['feed'];
    }
}
