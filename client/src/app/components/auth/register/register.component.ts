import { Component } from '@angular/core';
import { AuthService } from 'src/app/components/auth/auth.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  public pageData = "loading...";
  constructor(private _authService: AuthService) { }
  public ngOnInit() {
    this.getPageData()
  }
  private getPageData() {
    this._authService.getRegisterPageData().subscribe({
      next: (res) => {
        this.pageData = res;
      },
      error: err => {
        this.pageData = err.statusText;
      }
    })
  }
}
