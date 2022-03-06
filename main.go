/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	impsi "github.com/w6d-io/imps-injector/apis/secret/v1alpha1"
	sc "github.com/w6d-io/imps-injector/controllers/secret"
	"github.com/w6d-io/imps-injector/internal/config"
	"github.com/w6d-io/x/cmdx"
	"github.com/w6d-io/x/logx"
	"github.com/w6d-io/x/pflagx"
	//+kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
	Serve    = &cobra.Command{
		Use:   "serve",
		Short: "run operator",
		Run:   ServeFct,
	}
	Cmd = &cobra.Command{
		Use:   "help",
		Short: "Show this help",
		Run:   Help,
	}
	OsExit = os.Exit
	mgr    manager.Manager
)

func Help(cmd *cobra.Command, _ []string) {
	_ = cmd.Help()
}

func ServeFct(_ *cobra.Command, _ []string) {
	var err error
	mgr, err = ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     viper.GetString(config.ViperKeyMetricsListen),
		Port:                   9443,
		HealthProbeBindAddress: viper.GetString(config.ViperKeyProbeListen),
		LeaderElection:         viper.GetBool(config.ViperKeyLeader),
		LeaderElectionID:       "imps-injector.w6d.io",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		OsExit(1)
	}
	if err = (&sc.ImagePullSecretInjectorReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ImagePullSecretInjector")
		OsExit(1)
	}
	//+kubebuilder:scaffold:builder
	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		OsExit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		OsExit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		OsExit(1)
	}
}

func init() {
	cobra.OnInitialize(config.Init)

	pflagx.CallerSkip = -1
	pflagx.Init(Serve, &config.CfgFile)
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(impsi.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	Cmd.AddCommand(cmdx.Version(&config.Version, &config.Revision, &config.Built))
	Cmd.AddCommand(Serve)
	if err := Cmd.Execute(); err != nil {
		logx.WithName(context.Background(), "Main.Command").
			Error(err, "exec command failed")
		OsExit(1)
	}
}
