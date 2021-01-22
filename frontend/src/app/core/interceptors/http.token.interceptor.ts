import { Injectable, Injector } from '@angular/core';
import { HttpEvent, HttpInterceptor, HttpHandler, HttpRequest } from '@angular/common/http';
import { Observable } from 'rxjs';

import { JwtService } from '../services';
import { UserService } from '../services';
import { environment } from "../../../environments/environment";

@Injectable()
export class HttpTokenInterceptor implements HttpInterceptor {
  constructor(private jwtService: JwtService) {}

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const headersConfig = {
      'Content-Type': 'application/json',
      'Accept': 'application/json'
    };

    const token = this.jwtService.getToken();
    let tokenName = "Bearer";

    if (req.url === environment.api_url + "/products") tokenName = "Token";

    if (token) {
      headersConfig['Authorization'] = `${tokenName} ${token}`;
    }

    console.log(headersConfig);

    const request = req.clone({ setHeaders: headersConfig });
    return next.handle(request);
  }
}
