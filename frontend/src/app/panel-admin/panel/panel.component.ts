import { Component, Input, OnInit } from "@angular/core";
import { FormBuilder, FormControl, FormGroup } from "@angular/forms";
import { StatsService } from "../../core";

@Component({
  selector: "app-panel",
  templateUrl: "./panel.component.html",
  styles: [],
})
export class PanelComponent implements OnInit {
  @Input() currentUsers: any;
  productForm: FormGroup;
  commentMessage = new FormControl();
  productName = new FormControl();
  productPrice = new FormControl();
  productImage = new FormControl();
  productDescription = new FormControl();

  constructor(private fb: FormBuilder, private statsService: StatsService) {}

  isNewProduct: boolean = false;

  ngOnInit() {
    this.statsService.getCurrentUsersCache().subscribe((data) => {
      this.currentUsers = data.current_users
    });
    
    this.productForm = this.fb.group({
      productName: "",
      productPrice: "",
      productImage: "",
      productDescription: "",
    });
  }

  toogleNewProductForm(): void {
    console.log(this.currentUsers);

    this.isNewProduct = !this.isNewProduct;
  }

  submitProduct() {
    console.log(this.productForm.getRawValue());
  }
}
