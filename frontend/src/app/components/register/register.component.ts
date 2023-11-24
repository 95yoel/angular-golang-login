import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { UserForm, UserReg } from '../../interfaces/user-reg';
import { HttpService } from '../../services/http/http.service';
import { ValidationRegService } from '../../services/validation-reg/validation-reg.service';
import { AuthService } from '../../services/auth/auth.service';
import { RoutesHandlerService } from '../../services/routes-handler/routes-handler.service';


@Component({
  selector: 'app-register',
  standalone: true,
  imports: [CommonModule,RouterModule,FormsModule],
  templateUrl: './register.component.html',
  styleUrl: './register.component.css'
})
export class RegisterComponent {

  formData:UserForm={
    name: '',
    surname: '',
    email: '',
    password: '',
    confirmPassword: ''
  }

  constructor(private http:HttpService,private validator:ValidationRegService,private routesHandler:RoutesHandlerService){}


  ngOnInit(){
    this.routesHandler.verifyUserStatus();
  }

  onSubmit(){
    // Validate the form data using the validation service
    this.validator.validation(this.formData).then(
      ()=>{
        /*
        Create a user registration object userReg based on the UserForm interface
        With destructuring copy common properties marking confirmPassword as optional
        */
        const { confirmPassword, ...userReg }: UserReg & { confirmPassword?: string } = this.formData;
        
        // Call the HTTP service to register the user

        this.http.register(userReg).subscribe(
          (res)=>{
            console.log(res);
          }
        );
      }
    ).catch(
      // Log validation errors
      (error)=>{
        console.warn("Error : \n",error)
      }
    );
  }

}
