import { Component, OnInit, OnDestroy} from '@angular/core';
import { Subscription } from 'rxjs';
import { UsersService } from 'src/app/services/user/users.service';
import { TokenService } from "../../../services/token/token.service";
import { ChatService } from "src/app/services/socket/chat.service";

export interface MessageData {
  id : any;
  message : any;
}

@Component({
  selector: 'app-layout',
  templateUrl: './layout.component.html',
  styleUrls: ['./layout.component.scss']
})
export class LayoutComponent implements OnInit { 

  private readonly mediaWatcher: Subscription;
  constructor(private userService: UsersService, private tokenService: TokenService, public chat: ChatService) {
  }
  ngOnInit() { 

  }

  // ngOnDestroy(): void {
  //   this.mediaWatcher.unsubscribe();
  // }

}
