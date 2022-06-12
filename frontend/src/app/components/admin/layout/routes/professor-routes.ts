export const professorRoutes = [
  {
    path: 'dashboard',
    loadChildren: () =>
      import('../../dashboard/dashboard.module').then(m => m.DashboardModule),
    data: { icon: 'dashboard', text: 'Dashboard' }
  },
  {
    path: 'subject/:id',
    loadChildren: () =>
    import('../../dashboard/subject-dashboard/subject-dashboard.component').then(m => m.SubjectDashboardComponent),
    data: { text: 'Subject' }
  },
  {
    path: 'presentation',
    loadChildren: () =>
    import('../../dashboard/presentation-pdf/presentation-pdf.component').then(m => m.PresentationPdfComponent),
    data: { text: 'Presentation' }
  },
  {
    path: 'grades',
    loadChildren: () =>
    import('../../dashboard/grades/grades.component').then(m => m.GradesComponent),
    data: { text: 'Grades' }
  },
];
