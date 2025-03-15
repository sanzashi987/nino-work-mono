export const options: any = {
  type: 'CSuite',
  name: '背景图',
  children: {
    type: {
      type: 'CRadio',
      default: 'image',
      options: [
        { label: '图片', value: 'image' },
        { label: '渐变色', value: 'gradient' },
      ],
    },
    image: {
      type: 'CUpload',
      accept: 'image/*',
      showInPanel: {
        conditions: [['.type', '$=', 'image']],
      },
    },
    gradient: {
      type: 'CColor',
      themeEnable:false,
      modes: ['gradient'],
      showInPanel: {
        conditions: [['.type', '$=', 'gradient']],
      },
    },
  },
};
