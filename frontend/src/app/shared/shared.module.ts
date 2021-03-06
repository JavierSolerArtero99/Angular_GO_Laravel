import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { RouterModule } from '@angular/router';

import { ArticleListComponent, ArticleMetaComponent, ArticlePreviewComponent } from './article-helpers';
import { FavoriteButtonComponent, FollowButtonComponent } from './buttons';
import { ListErrorsComponent } from './list-errors.component';
import { ShowAuthedDirective } from './show-authed.directive';
import { ListProductsComponent } from '../product/list-products/list-products.component';
import { ProductPreviewComponent } from '../product/product-preview/product-preview.component';
import { ProductDetailsComponent } from '../product/product-details/product-details.component';
import { ProductCommentComponent } from '../product/product-comment/product-comment.component';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    RouterModule,
  ],
  declarations: [
    ProductCommentComponent,
    ListProductsComponent,
    ProductDetailsComponent,
    ProductPreviewComponent,
    ArticleListComponent,
    ArticleMetaComponent,
    ArticlePreviewComponent,
    FavoriteButtonComponent,
    FollowButtonComponent,
    ListErrorsComponent,
    ShowAuthedDirective,
  ],
  exports: [
    ProductCommentComponent,
    ListProductsComponent,
    ProductDetailsComponent,
    ProductPreviewComponent,
    ArticleListComponent,
    ArticleMetaComponent,
    ArticlePreviewComponent,
    CommonModule,
    FavoriteButtonComponent,
    FollowButtonComponent,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    ListErrorsComponent,
    RouterModule,
    ShowAuthedDirective,
  ],
})
export class SharedModule {}
