import { Routes } from '@angular/router';
import { RegisterComponent } from './components/register/register.component';
import { LoginComponent } from './components/login/login.component';
import { PrivateComponent } from './components/private/private.component';
import { guardRouterGuard } from './guard/guard-router.guard';
import { ErrorComponent } from './components/error/error.component';

export const routes: Routes = [
    /*
    "/" , "register" and "home"
    goes to the same RegisterComponent
    */
    {path:'',component:RegisterComponent,title:'Register'},
    {path:'register',component:RegisterComponent,title:'Register'},
    {path:'home',redirectTo: '', pathMatch: 'full',title:'Register'},
    {path:'login',component:LoginComponent,title:'Login'},
    /*
    Guard decides if the user is authorized to open private  
    */
    {path:'private',component:PrivateComponent,title:'Private',canActivate:[guardRouterGuard]},
    /*
    If the path is not in the list , go to ErrorComponent
    */
    {path:'**',component:ErrorComponent,title:'Error - Not Found'}
];
