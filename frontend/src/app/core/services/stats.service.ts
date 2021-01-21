import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";

import { ApiService } from "./api.service";
import { JwtService } from "./jwt.service";

@Injectable()
export class StatsService {
  constructor(
    private apiService: ApiService,
    private http: HttpClient,
    private jwtService: JwtService
  ) {}

  getCurrentUsersCache(): any {
    return this.apiService.get("/users/currentUsers");
  }

  getTotalUsersCache(): any {
    return this.apiService.get("/users/totalUsers");
  }

  getProductCache(): any {
    return this.apiService.get("/users/currentUsers");
  }
}
