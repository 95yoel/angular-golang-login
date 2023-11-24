import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { RoutesHandlerService } from './services/routes-handler/routes-handler.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  
  constructor(private routesHandler:RoutesHandlerService){}

  // Subscribe to router events when the app is initializated
  ngOnInit(){
    this.routesHandler.subscribe();
  }

  // Unsubscribe when the component is destroyed 
  ngOnDestroy(){
    this.routesHandler.unSubscribe();
  }


}
