import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { UserForm } from '../../interfaces/user-reg';


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

  onSubmit(){
    console.log(this.formData);
  }

}
