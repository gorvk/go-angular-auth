import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  public userData: string = "loading...";
  public URL = "";
  public buttonLabel = "create new account";
  constructor(private httpClient: HttpClient) { }

  public ngOnInit() {
    this.getAuthPageData("http://localhost:9090");
  }

  private getAuthPageData(url: string) {
    this.userData = "loading...";
    this.httpClient.post(url, "GK 👨‍💻", { responseType: "text" }).subscribe({
      next: (res) => {
        this.userData = res;
      },
      error: err => {
        this.userData = err.statusText;
      }
    })
  }

  switchAuth(event: Event) {
    event.preventDefault();
    if (this.buttonLabel === "login to your account") {
      this.buttonLabel = "create new account";
      this.getAuthPageData("http://localhost:9090");
    } else {
      this.buttonLabel = "login to your account";
      this.getAuthPageData("http://localhost:9090/register");
    }
  }
}