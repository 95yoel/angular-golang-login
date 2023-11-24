/*
    Interfaces for user register
*/

export interface UserReg {
    name: string;
    surname: string;
    email: string;
    password: string;
}

export interface UserForm{
    name: string;
    surname: string;
    email: string;
    password: string;
    confirmPassword:string;
}