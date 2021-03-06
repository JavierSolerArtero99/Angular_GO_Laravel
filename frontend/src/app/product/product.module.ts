import { NgModule } from "@angular/core";
import { CommonModule } from "@angular/common";

import { ProductRoutingModule } from "./product-routing.module";
import { ListProductsComponent } from "./list-products/list-products.component";
import { SharedModule } from "../shared";
import { MarkdownPipe } from "../article/markdown.pipe";
import { ProductDetailsComponent } from "./product-details/product-details.component";

@NgModule({
  imports: [ProductRoutingModule, SharedModule],
  declarations: [MarkdownPipe],
})
export class ProductModule {}
