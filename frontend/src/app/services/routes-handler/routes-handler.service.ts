import { Location } from '@angular/common';
import { Injectable } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { AuthService } from '../auth/auth.service';


@Injectable({
  providedIn: 'root'
})
export class RoutesHandlerService {

  private subscriber : Subscription = new Subscription;
  
  constructor(private location:Location,private router:Router,private auth:AuthService) { }

  // Subscribe to events emmitted by the router
  subscribe(){
    // This function detects events emitted by the router  
    this.subscriber = this.router.events.subscribe(
      (event)=>{
        // If the event emmited is a NavigationEnd event , url will change to "/"
        // This way the route is always localhost:xxxx/
        if(event instanceof NavigationEnd){
          this.location.replaceState("/");
        }
      }
    );
  }

  // Unsubscribe from router events and prevent memory leaks
  unSubscribe(){
    this.subscriber.unsubscribe();
  }


  /*
  When the user enters into a component this funcion verify is the user has credentials on the 
  localstorage , if its correct navigate to private
  */
  verifyUserStatus(){
    if(this.auth.isValidJWT()){
      this.router.navigate(['/private']);
    }
  }


}
