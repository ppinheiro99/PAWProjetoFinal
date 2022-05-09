import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TokenService } from "../../../../services/token/token.service";
import { UsersService } from "../../../../services/user/users.service";
@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  
  ngOnInit() {

  }

}


