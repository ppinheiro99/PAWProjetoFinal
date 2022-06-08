import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { PresentationPdfComponent } from './presentation-pdf/presentation-pdf.component';
import { SubjectDashboardComponent } from './subject-dashboard/subject-dashboard.component';

const routes: Routes = [
  {
    path: '',
    component: HomeComponent
  },
  {
    path: 'subject/:id',
    component: SubjectDashboardComponent
  },
  {
    path: 'presentation',
    component: PresentationPdfComponent
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DashboardRoutingModule { }
