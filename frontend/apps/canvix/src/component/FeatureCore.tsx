import React, { Component, createContext } from 'react';
import { noop, returnVoidObject } from '@nino-work/shared';
import type { EditorFeaturesRegisterType, EditorFeaturesType, FeatureItemProps, FeatureRuntimeMap } from '@/types';

export const EditorFeatures = createContext<EditorFeaturesType>({});
export const EditorFeaturesRegister = createContext<EditorFeaturesRegisterType>({
  registerFeatures: () => [],
  unregisterFeatures: noop,
  getFeaturesAsync: returnVoidObject
});

export const parseFeatures = (items: FeatureItemProps[]) => Object.fromEntries(
  items.map((item) => {
    const { id, shortcutNode } = item;
    return [id, { ...item, shortcutNode }];
  })
);

type FeatureCoreProps = {
  initFeatures: FeatureItemProps[];
  children: React.ReactNode;
};
type FeatureCoreState = {
  features: FeatureRuntimeMap;
};

class FeatureCore extends Component<FeatureCoreProps, FeatureCoreState> {
  registers;

  state = { features: {} };

  /** update the features and make it detectable immediately */
  featuresLastest: FeatureRuntimeMap = {};

  constructor(props: FeatureCore['props']) {
    super(props);
    this.featuresLastest = parseFeatures(props.initFeatures);
    this.state = { features: { ...this.featuresLastest } };
    this.registers = {
      registerFeatures: this.registerFeatures,
      unregisterFeatures: this.unregisterFeatures,
      getFeaturesAsync: this.getFeaturesAsync
    };
  }

  registerFeatures = (items: FeatureItemProps[]) => {
    const featurePatch = parseFeatures(items);
    this.featuresLastest = { ...this.featuresLastest, ...featurePatch };
    this.setState(() => ({ features: this.featuresLastest }));
    return items.map((item) => item.id);
  };

  unregisterFeatures = (ids: string[]) => {
    const nextFeature = { ...this.featuresLastest };
    ids.forEach((id) => {
      delete nextFeature[id];
    });
    this.featuresLastest = nextFeature;
    this.setState(() => ({ features: nextFeature }));
  };

  getFeaturesAsync = () => this.featuresLastest;

  render() {
    return (
      <EditorFeatures.Provider value={this.state.features}>
        <EditorFeaturesRegister.Provider value={this.registers}>
          {this.props.children}
        </EditorFeaturesRegister.Provider>
      </EditorFeatures.Provider>
    );
  }
}

export default FeatureCore;
