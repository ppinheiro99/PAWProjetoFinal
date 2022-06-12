import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { alunoRoutes } from '../layout/routes/aluno-routes';
import { professorRoutes } from '../layout/routes/professor-routes';
import { GradesComponent } from './grades/grades.component';
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
   {
     path: 'grades',
     component: GradesComponent
   },
    // ...professorRoutes,
    // ...alunoRoutes
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DashboardRoutingModule { }
