export const alunoRoutes = [
  {
    path: 'dashboard',
    loadChildren: () =>
      import('../../dashboard/dashboard.module').then(m => m.DashboardModule),
    data: { icon: 'dashboard', text: 'Dashboard' }
  },
  {
    path: 'grades',
    loadChildren: () =>
    import('../../dashboard/grades/grades.component').then(m => m.GradesComponent),
    data: { text: 'Grades' }
  },
];
