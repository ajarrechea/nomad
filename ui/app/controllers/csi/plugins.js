import Controller from '@ember/controller';

export default class PluginsController extends Controller {
  isForbidden = false;

  breadcrumbs = [
    {
      label: 'Storage',
      args: ['csi.index'],
    },
  ];
}
