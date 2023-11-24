import { CanActivateFn } from '@angular/router';

export const guardRouterGuard: CanActivateFn = (route, state) => {
  return true;
};
