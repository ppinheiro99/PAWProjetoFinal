import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatTableModule } from '@angular/material/table';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatDialogModule } from '@angular/material/dialog';
import { DashboardRoutingModule } from './dashboard-routing.module';
import { HomeComponent } from './home/home.component';
import { AddSubjectComponent } from './add-subject/add-subject.component';
import { FormsModule } from '@angular/forms';
import { SubjectDashboardComponent } from './subject-dashboard/subject-dashboard.component';
import { authInterceptorProviders } from 'src/app/helpers/auth.interceptor';
import { AddPresentationComponent } from './add-presentation/add-presentation.component';
import { MatSelectModule } from '@angular/material/select';
import { PresentationPdfComponent } from './presentation-pdf/presentation-pdf.component';
@NgModule({
  imports: [
    CommonModule,
    DashboardRoutingModule,
    MatGridListModule,
    FlexLayoutModule,
    MatCardModule,
    MatTableModule,
    MatButtonModule,
    MatIconModule,
    MatDialogModule,
    FormsModule,
    MatSelectModule
  ],
  declarations: [HomeComponent, AddSubjectComponent, SubjectDashboardComponent, AddPresentationComponent, PresentationPdfComponent],
  providers: [authInterceptorProviders]
})
export class DashboardModule {}
