import { Component, Input, OnInit } from "@angular/core";
import { FormBuilder, FormControl, FormGroup } from "@angular/forms";
import { StatsService } from "../../core/services/stats.service";

@Component({
  selector: "app-panel",
  templateUrl: "./panel.component.html",
  styles: [],
})
export class PanelComponent implements OnInit {
  currentUsers: number = 0;
  moneyEarned: number = 0;
  totalUsers: number = 0;
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
      this.currentUsers = data.users;
    });

    this.statsService.getTotalUsersCache().subscribe((data) => {
      this.totalUsers = data.users;
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
