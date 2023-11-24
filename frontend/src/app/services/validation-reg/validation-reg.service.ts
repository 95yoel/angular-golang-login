import { Injectable } from '@angular/core';
import { UserForm } from '../../interfaces/user-reg';

@Injectable({
  providedIn: 'root'
})
export class ValidationRegService {

  /*
  Throws an error if any of the validation funcitions encounter issues
  This funtion prevent the execution of unnecesary code
  */
  async validation(form:UserForm){
    this.validateEmpty(form);
    this.validatePassword(form);
    this.validateEmptyFields(form);
  }

  /*
  Throws an error is a field is empty or too short
  */
  validateEmptyFields(data:UserForm){
    if (data.name.length<3){
      throw new Error("Name too short")
    }
    if(data.surname.length<3){
      throw new Error("Surname too short")
    }
    if(!this.validateEmail(data.email)){
      throw new Error("Email is not formatted correctly")
    }
    if(data.password.length <8 || data.confirmPassword.length <8){
      throw new Error("Password must be at least 8 characters")
    }
    
  }
  /*
  It iterates through the properties of all the fields if it finds an empty one it throws an error
  */
  validateEmpty(data:UserForm){
    for (const key in data) {
      if (key in data && data[key as keyof UserForm] === '') {
        throw new Error("Empty values at form");
      }
    }
  }
  /*
  Validate email with a regex expression
  */
  validateEmail(email: string): boolean {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/; //Regex expression to test the email
    return emailRegex.test(email);
  }
  /*
  Throw error if passwords don`t match
  */
  validatePassword(data: UserForm){
    if (data.password !== data.confirmPassword) {
      throw new Error("Passwords do not match");
    }
  }



}
