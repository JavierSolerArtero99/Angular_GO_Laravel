import { Injectable } from "@angular/core";
import { environment } from "../../../environments/environment";
import { HttpHeaders, HttpClient, HttpParams } from "@angular/common/http";
import { Observable, throwError } from "rxjs";

import { JwtService } from "./jwt.service";
import { catchError } from "rxjs/operators";

@Injectable()
export class ApiService {
  constructor(
    private http: HttpClient,
    private goToken: JwtService,
    private laravelToken: JwtService
  ) {}

  private formatErrors(error: any) {
    return throwError(error.error);
  }

  private corsHeaders = new HttpHeaders({
    "Content-Type": "application/json",
    Accept: "application/json",
    "Access-Control-Allow-Origin": "*",
  });

  /* LARAVEL METHODS */

  get(path: string, params: HttpParams = new HttpParams()): Observable<any> {
    return this.http
      .get(`${environment.api_url}${path}`, { params })
      .pipe(catchError(this.formatErrors));
  }

  put(path: string, body: Object = {}): Observable<any> {
    return this.http
      .put(`${environment.api_url}${path}`, { body: JSON.stringify(body) })
      .pipe(catchError(this.formatErrors));
  }

  post(path: string, body: Object = {}): Observable<any> {
    return this.http
      .post(`${environment.api_url}${path}`, body)
      .pipe(catchError(this.formatErrors));
  }

  delete(path): Observable<any> {
    return this.http
      .delete(`${environment.api_url}${path}`)
      .pipe(catchError(this.formatErrors));
  }

  /* GO METHODS FOR: USERS*/

  getGo(path: string, params: HttpParams = new HttpParams()): Observable<any> {
    return this.http
      .get(`${environment.go_url}${path}`, { params })
      .pipe(catchError(this.formatErrors));
  }

  // PETITIONS - POST

  postGo(path: string, body: Object = {}): Observable<any> {
    console.log(`${environment.go_url}${path}`);

    return this.http
      .post(`${environment.go_url}${path}`, JSON.stringify(body))
      .pipe(catchError(this.formatErrors));
  }

  /* ----------------GO: PRODUCTS---------------- */

  // PETITIONS - GET

  getGoProducts(
    path: string,
    params: HttpParams = new HttpParams()
  ): Observable<any> {
    return this.http
      .get(`${environment.go_products_url}${path}`, { params })
      .pipe(catchError(this.formatErrors));
  }

  // PETITIONS - POST

  postGoProducts(path: string, body: Object = {}): Observable<any> {
    return this.http
      .post(`${environment.go_products_url}${path}`, JSON.stringify(body))
      .pipe(catchError(this.formatErrors));
  }

  deleteGoComments(path: string): Observable<any> {
    return this.http
      .delete(`${environment.go_products_url}${path}`)
      .pipe(catchError(this.formatErrors));
  }
}
