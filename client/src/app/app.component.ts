import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  public userData: User[] = []
  public URL = "http://localhost:3333/"
  // public URL = "http://localhost:3333/register"
  // public URL = "http://localhost:3333/nf"
  
  
  constructor(private httpClient: HttpClient){}
  
  public ngOnInit() {
    this.httpClient.get(this.URL).subscribe((res) => {
      const response = res as Response;
      this.userData = response.Data
    },
    err => {
      this.userData = err.error.Data      
    })
  }
}

interface User {
  Id: number, 
  Firstname: string, 
  Lastname: string
}

interface Response {
  Success: boolean,
  StatusCode: number,
  Data: any[]
}