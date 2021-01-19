import { Component, OnInit } from "@angular/core";
import { FormBuilder, FormControl, FormGroup } from "@angular/forms";

@Component({
  selector: "app-panel",
  templateUrl: "./panel.component.html",
  styles: [],
})
export class PanelComponent implements OnInit {
  productForm: FormGroup;
  commentMessage = new FormControl();
  productName = new FormControl();
  productPrice = new FormControl();
  productImage = new FormControl();
  productDescription = new FormControl();

  constructor(private fb: FormBuilder) {}

  isNewProduct: boolean = false;

  ngOnInit() {
    this.productForm = this.fb.group({
      productName: "",
      productPrice: "",
      productImage: "",
      productDescription: "",
    });
  }

  toogleNewProductForm(): void {
    this.isNewProduct = !this.isNewProduct;
  }

  submitProduct() {
    console.log(this.productForm.getRawValue());
  }
}
