import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../services/auth/auth.service';

export const guardRouterGuard: CanActivateFn = (route, state) => {

  const auth:AuthService = new AuthService();
  const router:Router = new Router;

  const isValid = auth.isValidJWT();

  if(isValid){
    return true;
  }else{
    router.navigate(["/"]);
    return false;
  }
};
