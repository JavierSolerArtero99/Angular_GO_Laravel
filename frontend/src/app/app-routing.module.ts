import { NgModule } from '@angular/core';
import { Routes, RouterModule, PreloadAllModules } from '@angular/router';
import { ListProductsComponent } from './product/list-products/list-products.component';

const routes: Routes = [
  {
    path: "settings",
    loadChildren: "./settings/settings.module#SettingsModule",
  },
  {
    path: "profile",
    loadChildren: "./profile/profile.module#ProfileModule",
  },
  {
    path: "editor",
    loadChildren: "./editor/editor.module#EditorModule",
  },
  {
    path: "article",
    loadChildren: "./article/article.module#ArticleModule",
  },
  {
    path: "product",
    loadChildren: "./product/product.module#ProductModule",
  },
  {
    path: "paneladmin",
    loadChildren: "./panel-admin/panel-admin.module#PanelAdminModule",
  },
];


@NgModule({
  imports: [
    RouterModule.forRoot(routes, {
      // preload all modules; optionally we could
      // implement a custom preloading strategy for just some
      // of the modules (PRs welcome 😉)
      preloadingStrategy: PreloadAllModules,
    }),
  ],
  // declarations: [ListProductsComponent],
  exports: [RouterModule],
})
export class AppRoutingModule {}
