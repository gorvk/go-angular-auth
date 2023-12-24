import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AuthService } from 'src/app/components/auth/auth.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {

  public registerForm: FormGroup;

  constructor(private _authService: AuthService) {
    this.registerForm = new FormGroup({
      username: new FormControl<string>("", [Validators.required, Validators.maxLength(8)]),
      password: new FormControl<string>("", [Validators.required, Validators.minLength(10), Validators.maxLength(10)]),
      age: new FormControl()
    })
  }
  public ngOnInit() { }

  public registerUser(event: SubmitEvent) {
    event.preventDefault()
    console.log(this.registerForm.valid);
    console.log(this.registerForm.value);
    this._authService.registerUser(this.registerForm.value).subscribe({
      next: (res) => {
        console.log(res);
      },
      error: err => {
        console.log(err);
      }
    })
  }
}
