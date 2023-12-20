import { Component } from '@angular/core';
import { AuthService } from 'src/app/components/auth/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  public pageData = "loading...";
  constructor(private _authService: AuthService) { }
  public ngOnInit() {
    this.getPageData()
  }
  private getPageData() {
    this._authService.getLoginPageData().subscribe({
      next: (res) => {
        this.pageData = res;
      },
      error: err => {
        this.pageData = err.statusText;
      }
    })
  }
}
