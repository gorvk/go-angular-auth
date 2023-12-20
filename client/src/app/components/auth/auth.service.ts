import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { urlConstant } from 'src/app/UrlConstants';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  constructor(private httpClient: HttpClient) { }
  
  public getLoginPageData(): Observable<string> {
    return this.httpClient.post(urlConstant.login, "GK 👨‍💻", { responseType: "text" })
  }

  public getRegisterPageData(): Observable<string> {
    return this.httpClient.post(urlConstant.register, "NK 👨‍💻", { responseType: "text" })
  }
}
