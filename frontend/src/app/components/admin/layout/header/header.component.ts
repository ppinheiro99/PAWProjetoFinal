import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Router } from '@angular/router';
import { TokenService } from "../../../../services/token/token.service";

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  @Output() sideNavToggled = new EventEmitter<void>();
  @Output() messageNavToggled = new EventEmitter<void>();
  name: string
  id:string
  isLoggedIn = false;
  displayedColumns = ['icon','message'];

  constructor(private readonly router: Router, private tokenService: TokenService) {}

  ngOnInit() {
    this.isLoggedIn = !!this.tokenService.getToken();
    if (this.isLoggedIn) {
      const user = this.tokenService.getUser();
      this.name = user.Name
      this.id = user.ID
    }
  }

  toggleSidebar() {
    this.sideNavToggled.emit()
  }

  toggleMessageBar(){
    this.messageNavToggled.emit()
  }

  onLoggedout() {
    localStorage.removeItem('isLoggedin')
    this.tokenService.signOut()
    this.router.navigate(['/login'])
    window.location.reload();
  }

  home(){
    this.router.navigate([''])
  }

  getGrades(){
    this.router.navigate(['grades'])
  }
}
